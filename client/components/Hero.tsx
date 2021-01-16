import React, { FC } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTwitter, faGithub } from '@fortawesome/free-brands-svg-icons';
import styles from './Hero.module.scss';
// import icon from '/images/icon.png';

const Hero: FC = () => (
    <div className={styles.hero}>
        <img alt="icon" src="/images/icon.png" />
        <div className={styles.links}>
            <a
                href="https://github.com/Le0tk0k"
                rel="noopener noreferrer"
                target="_blank"
                className={styles.github}
            >
                <FontAwesomeIcon icon={faGithub} />
                GitHub
            </a>
            <a
                href="https://github.com/Le0tk0k"
                rel="noopener noreferrer"
                target="_blank"
                className={styles.twitter}
            >
                <FontAwesomeIcon icon={faTwitter} />
                Twitter
            </a>
        </div>
    </div>
);

export default Hero;