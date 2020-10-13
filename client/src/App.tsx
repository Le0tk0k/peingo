import React, { FC } from 'react';
import { Switch, Route } from 'react-router';
import styles from './App.module.scss';
import Index from './pages/Index';
import QnA from './pages/QnA';

const App: FC = () => (
  <div className={styles.App}>
    <Switch>
      <Route exact path="/">
        <Index />
      </Route>
      <Route path="/:id">
        <QnA />
      </Route>
    </Switch>
  </div>
);
export default App;
