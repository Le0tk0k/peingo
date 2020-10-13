import React, { FC } from 'react';

interface Question {
  id: number;
  body: string;
}

const Questions: FC = () => {
  const questions: Question[] = [
    {
      id: 0,
      body: 'this is a question1.',
    },
    {
      id: 2,
      body: 'this is a question2.',
    },
  ];

  return (
    <ul>
      {questions.map((question) => (
        <li key={question.id}>{question.body}</li>
      ))}
    </ul>
  );
};

export default Questions;
