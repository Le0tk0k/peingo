import React, { FC, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getQuestion, GetQuestionRes } from '../api/client';

const Question: FC = () => {
  const { id }: any = useParams();

  const [qa, setQA] = useState({} as GetQuestionRes);

  useEffect(() => {
    const getQnA = async (req: number) => {
      const res: GetQuestionRes = await getQuestion(req);
      setQA(res);
    };

    try {
      getQnA(parseInt(id, 10));
    } catch (e) {
      console.error(e);
    }
  }, [id]);

  return (
    <>
      <p>{qa.id}</p>
      <p>{qa.question}</p>
      <p>{qa.answer}</p>
    </>
  );
};

export default Question;
