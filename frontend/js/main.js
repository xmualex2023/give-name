import { nameAPI } from './api.js';
import { i18n } from './i18n.js';

class NameGenerator {
    constructor() {
        this.form = document.getElementById('nameForm');
        this.input = document.getElementById('nameInput');
        this.resultContainer = document.getElementById('results');
        this.loadingIndicator = document.getElementById('loading');
        this.errorContainer = document.getElementById('error');
        this.langButtons = document.querySelectorAll('.lang-btn');
        
        this.init();
    }

    init() {
        // 绑定表单提交事件
        this.form.addEventListener('submit', (e) => this.handleSubmit(e));
        
        // 绑定输入验证
        this.input.addEventListener('input', () => this.validateInput());

        // 绑定语言切换事件
        this.langButtons.forEach(btn => {
            btn.addEventListener('click', () => {
                const lang = btn.getAttribute('data-lang');
                this.langButtons.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                
                const newLang = i18n.setLanguage(lang);
                // 如果有结果显示，重新渲染结果
                if (this.resultContainer.style.display !== 'none') {
                    this.rerenderResults();
                }
            });
        });

        // 初始化页面语言
        i18n.updatePageLanguage();
    }

    // 检测输入类型
    detectInputType(value) {
        const chinesePattern = /[\u4e00-\u9fa5]/;
        const englishPattern = /[a-zA-Z]/;
        
        if (chinesePattern.test(value)) {
            return 'chinese';
        } else if (englishPattern.test(value)) {
            return 'english';
        }
        return null;
    }

    // 验证输入
    validateInput() {
        const value = this.input.value.trim();
        const inputType = this.detectInputType(value);
        
        let isValid = false;
        if (inputType === 'chinese') {
            // 验证中文名（2-10个汉字）
            isValid = /^[\u4e00-\u9fa5]{2,10}$/.test(value);
        } else if (inputType === 'english') {
            // 验证英文名（2-50个字母，可包含空格和连字符）
            isValid = /^[a-zA-Z\s-]{2,50}$/.test(value);
        }
        
        this.input.classList.toggle('invalid', !isValid && value.length > 0);
        return isValid;
    }

    // 显示加载状态
    showLoading(show = true) {
        this.loadingIndicator.style.display = show ? 'block' : 'none';
        this.form.querySelector('button').disabled = show;
    }

    // 显示错误信息
    showError(message) {
        this.errorContainer.textContent = message;
        this.errorContainer.style.display = 'block';
        setTimeout(() => {
            this.errorContainer.style.display = 'none';
        }, 5000);
    }

    // 渲染单个名字建议
    renderSuggestion(suggestion, index) {
        if (!suggestion) return '';
        
        const currentLang = i18n.getCurrentLang();
        const titles = {
            zh: {
                meaning: '含义',
                cultural: '文化内涵',
                personality: '个性特征',
                english: '英文说明'
            },
            en: {
                meaning: 'Meaning',
                cultural: 'Cultural Significance',
                personality: 'Personality Traits',
                english: 'English Description'
            }
        };
        
        return `
            <div class="name-header">
                <h2 class="name-option">${suggestion.chinese_name || ''}</h2>
                <p class="pinyin">${suggestion.pinyin ? `${i18n.t('pinyin')}: ${suggestion.pinyin}` : ''}</p>
            </div>
            
            ${suggestion.characters && suggestion.characters.length > 0 ? `
                <div class="character-grid">
                    ${suggestion.characters.map(char => `
                        <div class="character-item">
                            <div class="character">${char.character || ''}</div>
                            <div class="char-pinyin">${char.pinyin || ''}</div>
                        </div>
                    `).join('')}
                </div>
            ` : ''}

            <div class="explanation">
                ${suggestion.meaning ? `
                    <div class="explanation-item">
                        <h3>${titles[currentLang].meaning}</h3>
                        <p>${suggestion.meaning}</p>
                    </div>
                ` : ''}
                
                ${suggestion.cultural_notes ? `
                    <div class="explanation-item">
                        <h3>${titles[currentLang].cultural}</h3>
                        <p>${suggestion.cultural_notes}</p>
                    </div>
                ` : ''}
                
                ${suggestion.personality ? `
                    <div class="explanation-item">
                        <h3>${titles[currentLang].personality}</h3>
                        <p>${suggestion.personality}</p>
                    </div>
                ` : ''}
                
                ${suggestion.english_intro ? `
                    <div class="explanation-item">
                        <h3>${titles[currentLang].english}</h3>
                        <p>${suggestion.english_intro}</p>
                    </div>
                ` : ''}
            </div>
        `;
    }

    // 重新渲染结果（用于语言切换）
    rerenderResults() {
        const suggestions = this.currentSuggestions;
        if (suggestions) {
            this.renderResults(suggestions);
        }
    }

    // 渲染所有结果
    renderResults(suggestions) {
        this.currentSuggestions = suggestions;
        
        // 先清空并设置容器为 flex 布局
        this.resultContainer.innerHTML = '';
        this.resultContainer.style.display = 'flex';
        this.resultContainer.style.gap = '1.5rem';
        this.resultContainer.style.justifyContent = 'center';
        this.resultContainer.style.flexWrap = 'nowrap';

        // 渲染每个建议
        suggestions.forEach((suggestion, index) => {
            const suggestionElement = document.createElement('div');
            suggestionElement.className = 'name-suggestion';
            suggestionElement.innerHTML = this.renderSuggestion(suggestion, index);
            this.resultContainer.appendChild(suggestionElement);
        });

        // 添加动画效果
        const cards = this.resultContainer.querySelectorAll('.name-suggestion');
        cards.forEach((card, index) => {
            setTimeout(() => {
                card.classList.add('visible');
            }, index * 100);
        });
    }

    // 处理表单提交
    async handleSubmit(e) {
        e.preventDefault();
        
        const value = this.input.value.trim();
        const inputType = this.detectInputType(value);
        
        if (!inputType) {
            this.showError(i18n.t('errorInvalidInput'));
            return;
        }

        if (!this.validateInput()) {
            const errorKey = inputType === 'english' ? 'errorInvalidNameChinese' : 'errorInvalidNameEnglish';
            this.showError(i18n.t(errorKey));
            return;
        }

        try {
            this.showLoading(true);
            this.resultContainer.style.display = 'none';
            
            const response = await nameAPI.getNameSuggestions(value, {
                language: i18n.getCurrentLang(),
                mode: inputType === 'english' ? 'chinese' : 'english'
            });
            
            this.renderResults(response.suggestions);
        } catch (error) {
            console.error('API Error:', error);
            this.showError(i18n.t('errorServer'));
        } finally {
            this.showLoading(false);
        }
    }
}

// 当 DOM 加载完成后初始化
document.addEventListener('DOMContentLoaded', () => {
    new NameGenerator();
}); 