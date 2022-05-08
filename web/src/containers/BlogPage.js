
import React from 'react';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import {withRouter} from './util';
import RetroMarkDownDisplay from '../components/common/RetroMarkdownDisplay';

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

  }
  componentDidMount() {
    let id = this.props.params.id
    fetch(`/api/blog/${id}`).then(res => {
      //todo: handle res code
      return res.json()
    }).then((res) => {
      this.setState({blog: res})
      console.log(res)
    }).catch(e => {
      //todo: handle err
      console.log(e)
    })
  }
  render() {
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}
        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <LeftSideBar></LeftSideBar>
          <RightBody>
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
}
export default withRouter(BlogPage)