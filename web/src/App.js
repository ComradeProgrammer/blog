import React from 'react';
import './App.css';
import {BrowserRouter, Routes, Route} from "react-router-dom";
import AllBlogsPage from './containers/AllBlogsPage';
import AboutMePage from './containers/AboutMePage';
import BlogPage from './containers/BlogPage';
import CategoryPage from './containers/CategoryPage';
import LoginPage from './containers/LoginPage';
import EditCategoryPage from './containers/EditCategoryPage';
import CreateBlogPage from './containers/CreateBlogPage';
import EditBlogPage from './containers/EditBlogPage';
import CreateCategoryPage from './containers/CreateCategoryPage';

class App extends React.Component {
  render() {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<AllBlogsPage />} />
          <Route path="aboutme" element={<AboutMePage />} />
          <Route path="blog" element={<AllBlogsPage />} />
          <Route path="blog/:id" element={<BlogPage />} />
          <Route path="category/:id" element={<CategoryPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/createblog" element={<CreateBlogPage />} />
          <Route path="/editblog/:id" element={<EditBlogPage />} />
          <Route path="/createcategory" element={<CreateCategoryPage />} />
          <Route path="/editcategory/:id" element={<EditCategoryPage />} />
        </Routes>
      </BrowserRouter>
    );
  }
}

export default App;