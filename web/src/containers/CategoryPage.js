import React from 'react';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import BlogBriefCard from '../components/BlogBriefCard';
import RetroButton from '../components/common/RetroButton';
import {isAdmin, withRouter} from './util';

class CategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      category: {}
    }
    this.onCreateBlogButtonClick = this.onCreateBlogButtonClick.bind(this)
    this.onEditCategotyButtonClick = this.onEditCategotyButtonClick.bind(this)
    this.onDeleteCategoryButtonClick=this.onDeleteCategoryButtonClick.bind(this)
  }
  componentDidMount() {
    this.fetchData()
  }

  fetchData() {
    let id = this.props.params.id
    fetch(`/api/category/${id}`).then(res => {
      if(!res.ok){
        window.alert("invalid blog")
        this.props.navigate(-1)
      }
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
    let hidden = !isAdmin()
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}
        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <LeftSideBar></LeftSideBar>
          <RightBody>
            <div style={{fontSize: "30px", width: "100%", textAlign: "center"}}>{this.state.category?.name}({this.state?.category?.blogs?.length})</div>
            <div style={{fontSize: "20px", width: "100%", textAlign: "center"}}>Description: {this.state.category?.description}</div>

            <div style={{fontSize: "15px", overflow: "hidden", whiteSpace: "nowrap", textOverflow: "ellipsis", width: "100%", textAlign: "center"}}>
              Create Time: {new Date(this.state.category?.createAt).toUTCString()}
              {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
              Last Modified: {new Date(this.state.category?.updateAt).toUTCString()}
            </div>

            <hr style={{border: "1px solid lightgreen"}} />
            <RetroButton hidden={hidden} style={{fontSize: "20px", margin: "5px"}} onClick={this.onCreateBlogButtonClick}>Create Blog</RetroButton>
            <RetroButton hidden={hidden} warning style={{fontSize: "20px", margin: "5px"}} onClick={this.onEditCategotyButtonClick}>Edit Category</RetroButton>
            <RetroButton hidden={hidden} danger style={{fontSize: "20px", margin: "5px"}} onClick={this.onDeleteCategoryButtonClick}>Delete Category</RetroButton>

            {blogItems}
          </RightBody>
        </div>
      </div>
    )
  }

  onCreateBlogButtonClick() {
    this.props.navigate(`/createblog`)
  }
  onEditCategotyButtonClick() {
    this.props.navigate(`/editcategory/${this.props.params.id}`)
  }
  onDeleteCategoryButtonClick() {
    let id = this.props.params.id
    let result = window.confirm("Do you really want to delete this category?")
    if (result) {
      fetch(`/api/category/${id}`,{
        method:"DELETE"
      }).then(async res => {
        //todo: handle res code
        if (res.ok) {
          alert("Operation succeeded")
          this.props.navigate(-1)
        } else {
          alert("Operation failed")
          let body = await res.json()
          console.log(body)
        }
      }).catch(e => {
        //todo: handle err
        console.log(e)
      })
    }
  }
}
export default withRouter(CategoryPage)