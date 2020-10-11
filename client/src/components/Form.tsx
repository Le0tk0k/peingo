import React, { FC } from 'react';
import styles from './Form.module.scss';

const Form: FC = () => (
  <div className={styles.form}>
    <p>気軽にどうぞ！</p>
    <textarea />
    <button>質問する</button>
  </div>
);

export default Form;
