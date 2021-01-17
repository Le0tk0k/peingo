import React from 'react';
import styles from '../styles/qna.module.scss';

export default function Question({qna}) {
    return (
        <div className={styles.wrapper}>
            <div className={styles.questionCard}>
                <div className={styles.questionCardBody}>
                    <p>{qna.question}</p>
                </div>
            </div>
            <div className={styles.answer}>{qna.answer}</div>
        </div>
    );
};

export const getStaticPaths = async () => {
    const res = await fetch("http://localhost:80/qnas")
    const data = await res.json()
    const qnas = data.qnas

    const paths = qnas.map((qna) => {
        return {
            params: {
                id: qna.id.toString(),
            },
        };
    });
    return {paths: paths, fallback: false}
}

export const getStaticProps = async ({params}) => {
    const res = await fetch(`http://localhost:80/qnas/${params.id}`)
    const qna = await res.json()

    return {
        props: {
            qna
        },
    }
}