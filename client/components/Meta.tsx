import React, {FC} from 'react';
import Head from 'next/head';

type Props = {
    content: string;
    url: string;
};

const Meta:FC<Props> = ({content, url}: Props) => {
    return (
        <Head>
            <title>
                peingo
            </title>
            <meta
                name="description"
                content="00の質問箱"
            />

            <meta name="format-detection" content="telephone=no" />
            <meta property="og:type" content="article" />
            <meta property="og:site_name" content="peingo" />
            <meta property="og:url" content={"https://peingo.vercel.app/" + url} />
            <meta
                property="og:title"
                content="00の質問箱"
            />
            <meta
                property="og:description"
                content="00の質問箱"
            />
            <meta
                property="og:image"
                content={"https://res.cloudinary.com/dyhhpykv4/image/upload/l_text:Sawarabi%20Gothic_80:" + content + ",w_1000,c_fit/v1611367657/peingo-ogp_ozaqel.png"}
            />

            <meta name='twitter:card' content='summary_large_image' />
            <meta
                name="twitter:title"
                content="00の質問箱"
            />
            <meta
                name='twitter:image'
                content={"https://res.cloudinary.com/dyhhpykv4/image/upload/l_text:Sawarabi%20Gothic_80:" + content + ",w_1000,c_fit/v1611367657/peingo-ogp_ozaqel.png"}
            />

        </Head>
    );
}

export default Meta;