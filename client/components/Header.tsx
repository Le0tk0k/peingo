import React, {FC} from 'react';
import Link from 'next/link';
import styles from './Header.module.scss';

const Header: FC = () => (
    <div className={styles.header}>
        <Link href='/'>
            <h1>00's Question Box</h1>
        </Link>
    </div>
);

export default Header;