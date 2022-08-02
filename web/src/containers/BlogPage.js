
import React from 'react';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import {withRouter, isAdmin, getUser} from './util';
import RetroMarkDownDisplay from '../components/common/RetroMarkdownDisplay';
import RetroButton from '../components/common/RetroButton';
import {Link} from "react-router-dom";
import RetroCard from '../components/common/RetroCard';
import RetroTextArea from '../components/common/RetroTextArea';


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
        comments: [],
      },
      id: null,
      comment: ""
    }
    this.onEditButtonClick = this.onEditButtonClick.bind(this)
    this.onDeleteBlogButtonClick = this.onDeleteBlogButtonClick.bind(this)
    this.onContentInputChange = this.onContentInputChange.bind(this)
    this.onPostCommentButtonClick = this.onPostCommentButtonClick.bind(this)

  }

  componentDidMount() {
    this.fetchData()
  }

  fetchData() {
    let id = this.props.params.id
    fetch(`/api/blog/${id}`).then(res => {
      //todo: handle res code
      if (!res.ok) {
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

  onContentInputChange(e) {
    this.setState({comment: e.target.value})
  }
  async onPostCommentButtonClick() {
    let user = getUser()
    if (user === null) {
      alert("You need to login first")
      return
    }

    let id = this.props.params.id
    let res = await fetch(`/api/comment`, {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        blogID: parseInt(id),
        userID: parseInt(user.ID),
        content: this.state.comment,
      })
    })
    if (res.ok) {
      alert("Operation Success")
      window.location.reload()
    } else {
      //todo: inform the problem
      let body = await res.json()
      console.log(body)
      alert("Operation failed")
    }
  }

  render() {
    let hidden = !isAdmin()
    let comments = []
    for (let i = 0; i < this.state.blog.comments.length; i++) {
      let obj = (
        <RetroCard style={{width: "100%", margin: "5px auto", padding: "3px"}} key={this.state.blog.comments[i].ID}>
          {this.state.blog.comments[i].user.userName} Posted at {this.state.blog.comments[i].createAt}
          <hr style={{border: "1px solid lightgreen"}} />
          {this.state.blog.comments[i].content}


        </RetroCard>
      )
      comments.push(obj)
    }

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

            <div style={{height: "25px"}}></div>
            <hr style={{border: "1px solid lightgreen"}} />
            <div style={{height: "25px"}}></div>

            {
              getUser() !== null ? (
                <RetroCard style={{width: "100%", margin: "0 auto", display: "block"}}>
                  <div style={{fontSize: "20px"}}>Want to share your opinions?</div>
                  <RetroTextArea style={{margin: "0 auto", width: "95%", fontSize: "15px", display: "block", height: "100px"}} value={this.state.comment} onChange={this.onContentInputChange} placeholder={"Post a comment here"} />
                  <RetroButton style={{fontSize: "20px", display: "block", width: "20%", marginLeft: "2%", marginTop: "3px"}} onClick={this.onPostCommentButtonClick}>Post Comment</RetroButton>
                </RetroCard>
              ) :
                (
                  <RetroCard style={{width: "100%", margin: "0 auto", display: "block"}}>
                    <div style={{fontSize: "20px"}}>Want to share your opinions? Please Log in</div>
                    <div style={{width: "20%", textAlign: "center", margin:"5px auto",border: "1px solid lightgreen"}}>
                      <Link to="/login" style={{color: "lightgreen", margin: "0 auto"}}>Login</Link>
                    </div>

                  </RetroCard>
                )
            }
            <div style={{height: "25px"}}></div>
            <hr style={{border: "1px solid lightgreen"}} />
            <div style={{height: "25px"}}></div>

            <div style={{fontSize: "20px"}}>Comments({this.state.blog.comments.length})</div>
            {comments}
          </RightBody>
        </div>
      </div>
    )
  }

  onDeleteBlogButtonClick() {
    let id = this.props.params.id
    let result = window.confirm("Do you really want to delete this blog?")
    if (result) {
      fetch(`/api/blog/${id}`, {
        method: "DELETE"
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