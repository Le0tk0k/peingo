import React, { FC, useEffect, useState } from 'react';
import QnA from '../entity/qna';
import { getQuestions, GetQuestionsRes } from '../api/client';

const Questions: FC = () => {
  const [data, setData] = useState({ qnas: [] as QnA[] });

  useEffect(() => {
    const getQnAs = async () => {
      const res: GetQuestionsRes = await getQuestions();
      setData({ qnas: res.qnas });
    };

    try {
      getQnAs();
    } catch (e) {
      console.log(e);
    }
  }, []);

  return (
    <ul>
      {data.qnas.map((qa: QnA) => (
        <li key={qa.id}>
          {qa.question},{qa.answer}
        </li>
      ))}
    </ul>
  );
};

export default Questions;
