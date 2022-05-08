import React from 'react';
import './App.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import AllBlogsPage from './containers/AllBlogsPage';
import AboutMePage from './containers/AboutMePage';
import BlogPage from './containers/BlogPage';
import CategoryPage from './containers/CategoryPage';

class App extends React.Component {
  render() {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<AllBlogsPage />} />
          <Route path="aboutme" element={<AboutMePage />} />
          <Route path="blog/:id" element={<BlogPage />} />
          <Route path="category/:id" element={<CategoryPage />} />
        </Routes>
      </BrowserRouter>
    );
  }
}

export default App;