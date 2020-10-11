import React, { FC } from 'react';
import styles from './App.module.scss';
import Index from './pages/Index';

const App: FC = () => (
  <div className={styles.App}>
    <Index />
  </div>
);

export default App;
