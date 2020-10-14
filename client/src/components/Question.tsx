import React, { FC, useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getQuestion, GetQuestionRes } from '../api/client';
import styles from './Question.module.scss';

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
    <div className={styles.wrapper}>
      <div className={styles.questionCard}>
        <div className={styles.questionCardBody}>{qa.question}</div>
      </div>
      <div className={styles.answer}>{qa.answer}</div>
    </div>
  );
};

export default Question;
