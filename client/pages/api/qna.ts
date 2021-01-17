import axios from 'axios';

const baseURL = 'http://localhost:80/';

const instance = axios.create({
    baseURL,
});

export const postQuestion = async (req: string) => {
    return await instance.post(`qnas`, { question: req });
};