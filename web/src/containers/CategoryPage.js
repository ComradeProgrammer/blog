import React from 'react';
import {withRouter} from './util';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import BlogBriefCard from '../components/BlogBriefCard';

class CategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      category: []
    }
  }
  componentDidMount() {
    this.fetchData()
  }

  fetchData() {
    let id = this.props.params.id
    fetch(`/api/category/${id}`).then(res => {
      //todo: handle res code
      return res.json()
    }).then((res) => {
      this.setState({category: res})
    }).catch(e => {
      //todo: handle err
      console.log(e)
    })
  }

  render() {
    let blogItems = []
    if (this.state.category?.blogs) {
      for (let i = 0; i < this.state.category.blogs.length; i++) {
        let enhancedBlog = {...this.state.category.blogs[i]}
        enhancedBlog.category = this.state.category
        blogItems.push((<BlogBriefCard blog={enhancedBlog} key={this.state.category.blogs[i].ID}></BlogBriefCard>))
      }
    }
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}
        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <LeftSideBar></LeftSideBar>
          <RightBody>
            <div style={{fontSize: "30px", width: "100%", textAlign: "center"}}>{this.state.category?.name}</div>
            <div style={{fontSize: "20px", width: "100%", textAlign: "center"}}>Description: {this.state.category?.description}</div>

            <div style={{fontSize: "15px", overflow: "hidden", whiteSpace: "nowrap", textOverflow: "ellipsis", width: "100%", textAlign: "center"}}>
              Create Time: {new Date(this.state.category?.createAt).toUTCString()}
              {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
              Last Modified: {new Date(this.state.category?.updateAt).toUTCString()}
            </div>

            <hr style={{border: "1px solid lightgreen"}} />
            {blogItems}
          </RightBody>
        </div>
      </div>
    )
  }
}
export default withRouter(CategoryPage)