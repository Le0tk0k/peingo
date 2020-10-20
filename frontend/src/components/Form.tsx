import React, { FC, useState } from 'react';
import styles from './Form.module.scss';

type Props = {
  onHandleSubmit: (question: string) => void;
};

const Form: FC<Props> = (props) => {
  const [question, setQuestion] = useState('');

  const handleChange = (e: any) => {
    setQuestion(e.target.value);
  };

  const handleClick = () => {
    props.onHandleSubmit(question);
    setQuestion('');
  };

  return (
    <div className={styles.form}>
      <p>気軽にどうぞ！</p>
      <textarea value={question} onChange={handleChange} />
      <button onClick={handleClick} disabled={!question}>
        質問する
      </button>
    </div>
  );
};

export default Form;
