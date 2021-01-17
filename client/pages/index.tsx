import React, {FC} from 'react';
import Link from 'next/link';
import Form from '../components/Form';
import Hero from '../components/Hero';
import QnA from "../entity/qna";
import {postQuestion} from './api/qna';
import Head from 'next/head'
import styles from '../styles/Home.module.scss'

const handleSubmit = async (q: string) => {
    try {
        const res = await postQuestion(q);
        console.log(res);
    } catch (e) {
        console.error(e);
    }
};

type Props = {
    qnas: QnA[],
}

const Home: FC<Props> = ({qnas}) => {
    return (
        <>
            <Hero/>
            <Form onHandleSubmit={handleSubmit}/>
            <div className={styles.questions}>
                <ul>
                    {qnas.map((qna: QnA) => (
                        <li key={qna.id}>
                            <Link href={`/${qna.id}`}>
                                <div>
                                    <div className={styles.questionCard}>
                                        <div className={styles.questionCardBody}>
                                            <p>{qna.question}</p>
                                        </div>
                                    </div>
                                    <div className={styles.answer}>{qna.answer}</div>
                                </div>
                            </Link>
                        </li>
                    ))}
                </ul>
            </div>
        </>
    )
};

export const getStaticProps = async () => {
    const res = await fetch("http://localhost:80/qnas")
    const data = await res.json()
    const qnas = data.qnas

    return {
        props: {
            qnas,
        },
    }
}

export default Home;