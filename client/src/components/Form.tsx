import React, { FC } from 'react';
import useState from 'react';
import styles from './Form.module.scss';

type Props = {
  onHandleSubmit: (question: string) => void;
};

const Form: FC<Props> = (onHandleSubmit) => {
  const [question, setQuestion] = useState('');

  const handleChange = (e: any) => {
    setQuestion(e.target.value);
  };

  const handleClick = () => {
    onHandleSubmit(question);
    setQuestion('');
  };

  return (
    <div className={styles.form}>
      <p>気軽にどうぞ！</p>
      <textarea value={question} onChange={handleChange} />
      <button onClick={handleClick}>質問する</button>
    </div>
  );
};

export default Form;
