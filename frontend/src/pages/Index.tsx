import React, { FC } from 'react';
import Header from '../components/Header';
import Hero from '../components/Hero';
import Form from '../components/Form';
import Questions from '../components/Questions';
import { postQuestion } from '../api/client';

const handleSubmit = async (q: string) => {
  try {
    const res = await postQuestion(q);
    console.log(res);
  } catch (e) {
    console.error(e);
  }
};

const Index: FC = () => (
  <>
    <Header />
    <Hero />
    <Form onHandleSubmit={handleSubmit} />
    <Questions />
  </>
);
export default Index;
