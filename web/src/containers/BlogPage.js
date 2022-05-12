
import React from 'react';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import {withRouter, isAdmin} from './util';
import RetroMarkDownDisplay from '../components/common/RetroMarkdownDisplay';
import RetroButton from '../components/common/RetroButton';

class BlogPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      blog: {
        title: null,
        ID: null,
        createAt: "",
        content: "",
        updateAt: "",
      },
      id: null,
    }
    this.onEditButtonClick = this.onEditButtonClick.bind(this)
    this.onDeleteBlogButtonClick = this.onDeleteBlogButtonClick.bind(this)
  }

  componentDidMount() {
    this.fetchData()
  }

  fetchData() {
    let id = this.props.params.id
    fetch(`/api/blog/${id}`).then(res => {
      //todo: handle res code
      if(!res.ok){
        window.alert("invalid blog")
        this.props.navigate(-1)
      }
      return res.json()
    }).then((res) => {
      this.setState({blog: res})
    }).catch(e => {
      //todo: handle err
      console.log(e)
    })
  }
  render() {
    let hidden = !isAdmin()
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}
        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <LeftSideBar></LeftSideBar>
          <RightBody>
            <RetroButton hidden={hidden} warning style={{fontSize: "20px", margin: "5px"}} onClick={this.onEditButtonClick}>Edit Blog</RetroButton>
            <RetroButton hidden={hidden} danger style={{fontSize: "20px", margin: "5px"}} onClick={this.onDeleteBlogButtonClick}>Delete Blog</RetroButton>

            <div style={{fontSize: "30px"}}>{this.state.blog.title}</div>
            <div style={{fontSize: "15px", overflow: "hidden", whiteSpace: "nowrap", textOverflow: "ellipsis"}}>
              Category: {this.state.blog?.category?.name}
              {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
              Create Time: {new Date(this.state.blog?.createAt).toUTCString()}
              {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
              Last Modified: {new Date(this.state.blog?.updateAt).toUTCString()}
            </div>

            <hr style={{border: "1px solid lightgreen"}} />
            <RetroMarkDownDisplay markdown={this.state.blog?.content}></RetroMarkDownDisplay>
          </RightBody>
        </div>
      </div>
    )
  }

  onDeleteBlogButtonClick() {
    let id = this.props.params.id
    let result = window.confirm("Do you really want to delete this blog?")
    if (result) {
      fetch(`/api/blog/${id}`,{
        method:"DELETE"
      }).then(async res => {
        //todo: handle res code
        if (res.ok) {
          alert("Operation succeeded")
          this.props.navigate(-1)
        } else {
          let body = await res.json()
          console.log(body)
        }
      }).catch(e => {
        //todo: handle err
        console.log(e)
      })
    }
  }

  onEditButtonClick() {
    this.props.navigate(`/editblog/${this.props.params.id}`)
  }
}
export default withRouter(BlogPage)