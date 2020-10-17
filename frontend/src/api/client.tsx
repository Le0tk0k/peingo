import axios from 'axios';
import QnA from '../entity/qna';

const baseURL = 'http://localhost:1323/';

const instance = axios.create({
  baseURL,
});

export interface GetQuestionsRes {
  qnas: QnA[];
}

export interface GetQuestionRes {
  id: number;
  question: string;
  answer: string;
}

export const postQuestion = async (req: string) => {
  return await instance.post(`qnas`, { question: req });
};

export const getQuestions = async () => {
  const res = await instance.get<GetQuestionsRes>('qnas');

  return res.data;
};

export const getQuestion = async (req: number) => {
  const res = await instance.get<GetQuestionRes>(`qnas/${req}`);

  return res.data;
};