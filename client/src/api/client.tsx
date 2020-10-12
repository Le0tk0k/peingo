import axios from 'axios';

const baseURL = 'http://localhost:1323/';

const instance = axios.create({
  baseURL,
});

export const postQuestion = async (req: string) => {
  return await instance.post(`qnas`, { question: req });
};
