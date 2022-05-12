import '../App.css';
import React from 'react';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import RetroInput from '../components/common/RetroInput';
import RetroButton from '../components/common/RetroButton';
import {withRouter} from './util';
import RetroCard from '../components/common/RetroCard';


class CreateCategoryPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      height: 500,
      name: "",
      description: "",
      alertMsg: []
    }
    this.onDescriptionInputChange = this.onDescriptionInputChange.bind(this)
    this.onNameInputChange = this.onNameInputChange.bind(this)
    this.onCreateClick = this.onCreateClick.bind(this)
  }

  componentDidMount() {
    this.setState({height: window.innerHeight - 150})
  }

  render() {
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}

        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <LeftSideBar></LeftSideBar>
          <RightBody style={{height: this.state.height}}>
            <div style={{fontSize: "30px", width: "100%", textAlign: "center"}}>Create a new category</div>
            <hr style={{border: "1px solid lightgreen"}} />
            {this.state.alertMsg}
            <span style={{fontSize: "20px", marginTop: "20px", display: "inline-block"}}>Name</span><br />
            <RetroInput style={{margin: "0 auto", width: "95%", fontSize: "15px", display: "block", border: "2px solid lightgreen"}} onChange={this.onNameInputChange} />

            <span style={{fontSize: "20px", marginTop: "20px", display: "inline-block"}}>Description</span><br />
            <RetroInput style={{margin: "0 auto", width: "95%", fontSize: "15px", display: "block", border: "2px solid lightgreen"}} onChange={this.onDescriptionInputChange} />

            <RetroButton style={{fontSize: "20px", margin: "20px auto", display: "block", width: "20%"}} onClick={this.onCreateClick}>Create</RetroButton>
          </RightBody>
        </div>
      </div>
    )
  }
  onNameInputChange(e) {
    this.setState({name: e.target.value})
  }

  onDescriptionInputChange(e) {
    this.setState({description: e.target.value})
  }

  onCreateClick() {
    let alertMsg = []
    if (this.state.name === "") {
      alertMsg.push((<RetroCard danger style={{display: "block", width: "100%", margin: "5px"}} key={1}>Name must not be empty</RetroCard>))
    }
    if (this.state.description === "") {
      alertMsg.push((<RetroCard danger style={{display: "block", width: "100%", margin: "5px"}} key={1}>Description must not be empty</RetroCard>))
    }
    if (alertMsg.length !== 0) {
      this.setState({alertMsg: alertMsg})
      return
    }
    fetch("/api/category", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        name: this.state.name,
        description: this.state.description,
      })
    }).then(async res => {
      if (res.ok) {
        //todo: inform the operation is successful
        this.props.navigate(-1)
      } else {
        let body = await res.json()
        console.log(body)
        alertMsg.push((<RetroCard danger style={{display: "block", width: "100%", margin: "5px"}} key={4}>body.error</RetroCard>))
        this.setState({alertMsg: alertMsg})
      }
    })

  }

}
export default withRouter(CreateCategoryPage)