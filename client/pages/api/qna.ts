import axios from 'axios';

export const baseURL = process.env.NEXT_APP_API_BASE_URL || 'http://localhost:80/';

const instance = axios.create({
    baseURL,
});

export const postQuestion = async (req: string) => {
    return await instance.post(`qnas`, { question: req });
};