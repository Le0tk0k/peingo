import React, { FC } from 'react';
import styles from './Hero.module.scss';
import icon from '../icon.png';

const Hero: FC = () => (
  <div className={styles.hero}>
    <img alt="icon" src={icon} />
    <div className={styles.link}>
      <a
        href="https://github.com/Le0tk0k"
        rel="noopener noreferrer"
        target="_blank"
      >
        GitHub
      </a>
      <a
        href="https://github.com/Le0tk0k"
        rel="noopener noreferrer"
        target="_blank"
      >
        Twitter
      </a>
    </div>
  </div>
);

export default Hero;
