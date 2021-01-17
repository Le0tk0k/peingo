import React, {FC} from 'react';
import Link from 'next/link';
import Header from '../components/Header';
import Form from '../components/Form';
import Hero from '../components/Hero';
import QnA from "../entity/qna";
import { postQuestion } from './api/qna';
import Head from 'next/head'
import styles from '../styles/Home.module.css'

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
            <Header/>
            <Hero/>
            <Form onHandleSubmit={handleSubmit} />
            <ul>
                {qnas.map((qna) => (
                    <li key={qna.id}>
                        <Link href={`/${qna.id}`}>
                            <a>{qna.question}</a>
                        </Link>
                    </li>
                ))}
            </ul>
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