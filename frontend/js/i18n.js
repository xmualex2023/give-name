// 语言配置
const translations = {
    en: {
        brandName: 'Elegance Names',
        title: 'Name Generator',
        subtitle: 'Enter any name to get its cultural translation',
        namePlaceholder: 'Enter your name (Chinese or English)',
        generateBtn: 'Generate Names',
        loading: 'Generating your names...',
        errorInvalidInput: 'Please enter a valid name in Chinese or English',
        errorInvalidNameChinese: 'Please enter a valid English name (2-50 letters)',
        errorInvalidNameEnglish: 'Please enter a valid Chinese name (2-10 characters)',
        errorServer: 'Server error. Please try again later.',
        pinyin: 'Pinyin',
        pronunciation: 'Pronunciation',
        culturalSignificance: 'Cultural Significance',
        personalityTraits: 'Personality Traits',
        meaningDescription: 'Meaning',
        englishIntro: 'English Introduction',
        
        // Footer translations
        footerAbout: 'About',
        footerDescription: 'Discover your perfect name that reflects your identity using advanced AI technology.',
        footerContact: 'Contact',
        footerSupport: 'Support',
        footerPrivacy: 'Privacy Policy',
        footerTerms: 'Terms of Service',
        footerFAQ: 'FAQ',
        footerCopyright: '© 2024 Name Generator - All rights reserved',
        footerPowered: 'Powered by AI Technology'
    },
    zh: {
        brandName: '雅名阁',
        title: '姓名生成器',
        subtitle: '输入任意名字，获取完美的文化翻译',
        namePlaceholder: '输入您的名字（中文或英文）',
        generateBtn: '生成名字',
        loading: '正在生成您的名字...',
        errorInvalidInput: '请输入有效的中文或英文名字',
        errorInvalidNameChinese: '请输入有效的英文名（2-50个字母）',
        errorInvalidNameEnglish: '请输入有效的中文名（2-10个汉字）',
        errorServer: '服务器错误，请稍后重试。',
        pinyin: '拼音',
        pronunciation: '发音',
        culturalSignificance: '文化内涵',
        personalityTraits: '个性特征',
        meaningDescription: '含义',
        englishIntro: '英文说明',
        
        // Footer translations
        footerAbout: '关于我们',
        footerDescription: '使用先进的 AI 技术，为您找到最适合的名字。',
        footerContact: '联系方式',
        footerSupport: '支持',
        footerPrivacy: '隐私政策',
        footerTerms: '服务条款',
        footerFAQ: '常见问题',
        footerCopyright: '© 2024 姓名生成器 - 保留所有权利',
        footerPowered: '由 AI 技术驱动'
    }
};

class I18n {
    constructor() {
        this.currentLang = 'zh';  // 默认使用中文
        this.translations = translations;
    }

    // 设置语言
    setLanguage(lang) {
        if (lang !== 'en' && lang !== 'zh') {
            console.error('不支持的语言:', lang);
            return this.currentLang;
        }
        this.currentLang = lang;
        this.updatePageLanguage();
        return this.currentLang;
    }

    // 切换语言
    toggleLanguage() {
        return this.setLanguage(this.currentLang === 'en' ? 'zh' : 'en');
    }

    // 获取当前语言
    getCurrentLang() {
        return this.currentLang;
    }

    // 获取翻译文本
    t(key) {
        return this.translations[this.currentLang][key] || key;
    }

    // 更新页面语言
    updatePageLanguage() {
        // 更新 HTML lang 属性
        document.documentElement.lang = this.currentLang;

        try {
            // 更新所有带有 data-i18n 属性的元素
            document.querySelectorAll('[data-i18n]').forEach(element => {
                const key = element.getAttribute('data-i18n');
                element.textContent = this.t(key);
            });

            // 更新所有带有 data-i18n-placeholder 属性的元素
            document.querySelectorAll('[data-i18n-placeholder]').forEach(element => {
                const key = element.getAttribute('data-i18n-placeholder');
                element.placeholder = this.t(key);
            });

            // 更新语言切换按钮文本
            const langBtn = document.getElementById('langToggle');
            if (langBtn) {
                langBtn.textContent = this.currentLang === 'en' ? '中文' : 'English';
            }
        } catch (error) {
            console.error('Error updating page language:', error);
        }
    }
}

export const i18n = new I18n(); 