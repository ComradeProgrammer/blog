import React from 'react';
import BlogBriefCard from '../components/BlogBriefCard';
import RetroButton from '../components/common/RetroButton';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import {isAdmin, withRouter} from './util';
class AllBlogsPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      blogs: []
    }
    this.onCreateBlogButtonClick = this.onCreateBlogButtonClick.bind(this)
    this.onCreateCategoryButtonClick = this.onCreateCategoryButtonClick.bind(this)
  }
  componentDidMount() {
    this.fetchData()
  }

  fetchData() {
    fetch("/api/blog").then(res => {
      //todo: handle res code
      return res.json()
    }).then((res) => {
      this.setState({blogs: res})
    }).catch(e => {
      //todo: handle err
      console.log(e)
    })
  }

  render() {
    let hidden = !isAdmin()
    let blogItems = []
    for (let i = 0; i < this.state.blogs.length; i++) {
      blogItems.push((<BlogBriefCard blog={this.state.blogs[i]} key={this.state.blogs[i].ID}></BlogBriefCard>))
    }
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}
        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <LeftSideBar></LeftSideBar>
          <RightBody>
            <div style={{fontSize: "30px", width: "100%", textAlign: "center"}}>All blogs({this.state.blogs.length})</div>
            <hr style={{border: "1px solid lightgreen"}} />
            <RetroButton hidden={hidden} style={{fontSize: "20px", margin: "10px"}} onClick={this.onCreateCategoryButtonClick}>Create category</RetroButton>
            <RetroButton hidden={hidden} style={{fontSize: "20px", margin: "5px"}} onClick={this.onCreateBlogButtonClick}>Create Blog</RetroButton>
            {blogItems}
          </RightBody>
        </div>
      </div>
    )
  }
  onCreateBlogButtonClick() {
    this.props.navigate(`/createblog`)
  }
  onCreateCategoryButtonClick() {
    this.props.navigate(`/createcategory`)
  }
}
export default withRouter(AllBlogsPage)