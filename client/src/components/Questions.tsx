import React, { FC, useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import QnA from '../entity/qna';
import { getQuestions, GetQuestionsRes } from '../api/client';
import styles from './Questions.module.scss';

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
        <Link to={`/${qa.id}`} key={qa.id}>
          <li>
            <div className={styles.questionCard}>
              <div className={styles.questionCardBody}>{qa.question}</div>
            </div>
            <div className={styles.answer}>{qa.answer}</div>
          </li>
        </Link>
      ))}
    </ul>
  );
};

export default Questions;
