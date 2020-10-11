import React, { FC } from 'react';
import Header from '../components/Header';
import Hero from '../components/Hero';
import Form from '../components/Form';
import Questions from '../components/Questions';

const handleSubmit = () => {};

const Index: FC = () => (
  <>
    <Header />
    <Hero />
    <Form onHandleSubmit={handleSubmit} />
    <Questions />
  </>
);
export default Index;
