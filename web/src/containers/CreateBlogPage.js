import '../App.css';
import React from 'react';
import Header from '../components/Header';
import RightBody from '../components/RightBody';
import {withRouter} from './util';
import RetroTextArea from '../components/common/RetroTextArea';
import RetroInput from '../components/common/RetroInput';
import RetroButton from '../components/common/RetroButton';
import RetroOption from '../components/common/RetroOption';
import RetroCard from '../components/common/RetroCard';
class CreateBlogPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      height: 500,
      categories: [],
      title: "",
      content: "",
      categoryID: null,
      alertMsg: []
    }
    this.onCategorySelectChange = this.onCategorySelectChange.bind(this)
    this.onContentInputChange = this.onContentInputChange.bind(this)
    this.onTitleInputChange = this.onTitleInputChange.bind(this)
    this.onConfirmClick = this.onConfirmClick.bind(this)
  }
  componentDidMount() {
    this.setState({height: window.innerHeight - 200})
    fetch("/api/category").then(res => {
      //todo: handle res code
      return res.json()
    }).then((res) => {
      let newCategories = []
      newCategories.push({ID: 0, name: "Please select a category"})
      newCategories.push(...res)
      this.setState({categories: newCategories})
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

          <RightBody style={{display: "block", marginLeft: "12%", padding: "30px"}}>

            <div style={{fontSize: "30px", width: "100%", textAlign: "center"}}>Post a new blog</div>
            <hr style={{border: "1px solid lightgreen"}} />
            {this.state.alertMsg}
            <span style={{fontSize: "20px"}}>Title</span><br />
            <RetroInput style={{margin: "0 auto", width: "100%", fontSize: "15px", display: "block", border: "2px solid lightgreen"}} onChange={this.onTitleInputChange} />

            <span style={{fontSize: "20px"}}>Category</span><br />
            <RetroOption data={this.state.categories} style={{width: "50%", fontSize: "15px", display: "block"}} onChange={this.onCategorySelectChange}></RetroOption>

            <span style={{fontSize: "20px"}}>Content</span><br />
            <RetroTextArea style={{margin: "0 auto", width: "100%", height: this.state.height + "px", fontSize: "15px", display: "block", minHeight: "300px"}} onChange={this.onContentInputChange} />

            <RetroButton style={{fontSize: "20px", margin: "20px auto", display: "block", width: "20%"}} onClick={this.onConfirmClick}>Confirm</RetroButton>
          </RightBody>
        </div>
      </div>
    )
  }

  onTitleInputChange(e) {
    this.setState({title: e.target.value})
  }

  onContentInputChange(e) {
    this.setState({content: e.target.value})
  }

  onCategorySelectChange(e) {
    this.setState({categoryID: e.target.value})
  }

  async onConfirmClick() {
    let alertMsg = []
    if (this.state.title === "") {
      alertMsg.push((<RetroCard danger style={{display: "block", width: "100%", margin: "5px"}} key={1}>Title must not be empty</RetroCard>))
    }
    if (this.state.categoryID === null || Number(this.state.categoryID)===0) {
      alertMsg.push((<RetroCard danger style={{display: "block", width: "100%", margin: "5px"}} key={2}>You must choose a category</RetroCard>))
    }
    if (this.state.content === "") {
      alertMsg.push((<RetroCard danger style={{display: "block", width: "100%", margin: "5px"}} key={3}>Content must not be empty</RetroCard>))
    }
    if (alertMsg.length !== 0) {
      this.setState({alertMsg: alertMsg})
      return
    }


    let res = await fetch("/api/blog", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        title: this.state.title,
        content: this.state.content,
        categoryID: this.state.categoryID
      })
    })
    if (res.ok) {
      //todo: inform the operation is successful
      this.props.navigate(-1)
    } else {
      //todo: inform the problem
      let body = await res.json()
      console.log(body)
      alertMsg.push((<RetroCard danger style={{display: "block", width: "100%", margin: "5px"}} key={4}>body.error</RetroCard>))
      this.setState({alertMsg: alertMsg})
    }
  }

}

export default withRouter(CreateBlogPage)