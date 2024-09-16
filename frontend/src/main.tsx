import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { HashRouter, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Overlay from './pages/Overlay';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <HashRouter>
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/overlay' element={<Overlay />} />
      </Routes>
    </HashRouter>
  </React.StrictMode>
);
