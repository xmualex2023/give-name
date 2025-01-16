// API 配置
const API_CONFIG = {
    baseURL: 'http://localhost:8080/api',
    timeout: 20000, // 20 秒超时
    headers: {
        'Content-Type': 'application/json'
    }
};

/**
 * 生成中文名字
 * @param {string} englishName - 英文名
 * @param {string} [language='en'] - 语言选项
 * @returns {Promise<Object>} 名字建议列表
 */
async function generateNames(englishName, language = 'en') {
    try {
        const response = await fetch(`${API_CONFIG.baseURL}/generate`, {
            method: 'POST',
            headers: API_CONFIG.headers,
            body: JSON.stringify({
                english_name: englishName,
                language: language
            })
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || '生成名字失败');
        }

        return await response.json();
    } catch (error) {
        console.error('API Error:', error);
        throw new Error(error.message || '服务器错误');
    }
}

/**
 * 格式化拼音（首字母大写）
 * @param {string} pinyin - 原始拼音
 * @returns {string} 格式化后的拼音
 */
function formatPinyin(pinyin) {
    return pinyin
        .split(' ')
        .map(word => word.charAt(0).toUpperCase() + word.slice(1))
        .join(' ');
}

/**
 * 格式化名字建议
 * @param {Object} suggestion - 名字建议对象
 * @returns {Object} 格式化后的名字建议
 */
function formatNameSuggestion(suggestion) {
    return {
        ...suggestion,
        pinyin: formatPinyin(suggestion.pinyin),
        characters: suggestion.characters.map(char => ({
            ...char,
            pinyin: formatPinyin(char.pinyin)
        }))
    };
}

// 导出 API 函数
export const nameAPI = {
    /**
     * 获取名字建议
     * @param {string} englishName - 英文名
     * @param {Object} options - 可选参数
     * @returns {Promise<Object>} 格式化后的名字建议
     */
    async getNameSuggestions(englishName, options = {}) {
        try {
            // 参数验证
            if (!englishName || typeof englishName !== 'string') {
                throw new Error('请输入有效的英文名');
            }

            // 调用 API
            const response = await generateNames(englishName, options.language);

            // 格式化响应数据
            return {
                ...response,
                suggestions: response.suggestions.map(formatNameSuggestion)
            };
        } catch (error) {
            console.error('Name Generation Error:', error);
            throw error;
        }
    }
}; 