async function generateNames() {
    const firstName = document.getElementById('firstName').value;
    const lastName = document.getElementById('lastName').value;
    
    try {
        const response = await fetch('http://localhost:8000/api/generate', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                firstName,
                lastName
            })
        });
        
        const data = await response.json();
        displayResults(data.names);
    } catch (error) {
        console.error('Error:', error);
        alert('Failed to generate names. Please try again.');
    }
}

function displayResults(names) {
    const container = document.querySelector('.name-cards-container');
    container.innerHTML = '';
    
    names.forEach(name => {
        const card = createNameCard(name);
        container.appendChild(card);
    });
    
    document.querySelector('.results-section').classList.remove('hidden');
}

function createNameCard(name) {
    // 创建名字卡片的DOM结构
    // ...
}

// 语言切换功能
const langButtons = document.querySelectorAll('.lang-btn');
const translations = {
    en: {
        title: 'Discover Your Perfect Chinese Name',
        subtitle: 'Experience the beauty of Chinese culture through a meaningful name that resonates with your identity',
        placeholder: 'Enter your full name (e.g. Barack Hussein Obama)',
        generateBtn: 'Generate Names',
        // ... 其他翻译内容
    },
    zh: {
        title: '探索您的完美中文名字',
        subtitle: '体验与您身份相呼应的富有文化内涵的中文名字',
        placeholder: '请输入您的英文名 (例如: Barack Hussein Obama)',
        generateBtn: '生成名字',
        // ... 其他翻译内容
    }
};

langButtons.forEach(btn => {
    btn.addEventListener('click', () => {
        // 移除所有按钮的激活状态
        langButtons.forEach(b => b.classList.remove('active'));
        // 激活当前按钮
        btn.classList.add('active');
        // 切换语言
        const lang = btn.dataset.lang;
        changeLang(lang);
    });
});

function changeLang(lang) {
    const t = translations[lang];
    document.querySelector('.header h1').textContent = t.title;
    document.querySelector('.subtitle').textContent = t.subtitle;
    document.querySelector('#englishName').placeholder = t.placeholder;
    document.querySelector('#generateBtn').textContent = t.generateBtn;
    // ... 更新其他元素的文本
} 