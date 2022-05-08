import React from 'react';
import RetroCard from './common/RetroCard';

class BlogBriefCard extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      blog: props.blog,
      style: {
        width: "98%",
        padding: "3px",
        margin: "5px auto",
        height: "100px",
        display: "block",
        overflow: "hidden",
        whiteSpace: "nowrap",
        textOverflow: "ellipsis"
      }
    }
    //copy the remaining style
    for (let key in props.style) {
      this.state.style[key] = props.style[key]
    }
  }

  render() {
    return (
      <RetroCard style={this.state.style}>
        <div style={{fontSize: "30px", overflow: "hidden", whiteSpace: "nowrap", textOverflow: "ellipsis"}}>{this.state.blog.title}</div>
        <div style={{fontSize: "15px", overflow: "hidden", whiteSpace: "nowrap", textOverflow: "ellipsis"}}>
          Category: {this.state.blog.category.name}
          {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
          Create Time: {new Date(this.state.blog.createAt).toUTCString()}
          {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
          Last Modified: {new Date(this.state.blog.updateAt).toUTCString()}
        </div>
        <div style={{fontSize: "20px", overflow: "hidden", whiteSpace: "nowrap", textOverflow: "ellipsis"}}>{this.state.blog.content}</div>
      </RetroCard>
    )
  }
}
export default BlogBriefCard