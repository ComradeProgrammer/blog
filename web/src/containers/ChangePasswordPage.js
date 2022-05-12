import '../App.css';
import React from 'react';
import {withRouter} from './util';
import Header from '../components/Header';
import RetroCard from '../components/common/RetroCard';
import RetroInput from '../components/common/RetroInput';
import RetroButton from '../components/common/RetroButton';

class ChangePassword extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      userName: "",
      password: "",
      newPassword: "",
      alertMsg: "",
    }
    this.onPasswordChange = this.onPasswordChange.bind(this)
    this.onUserNameChange = this.onUserNameChange.bind(this)
    this.onConfirmButtonClick = this.onConfirmButtonClick.bind(this)
    this.onNewPasswordChange = this.onNewPasswordChange.bind(this)
  }

  onUserNameChange(e) {
    this.setState({userName: e.target.value})
  }
  onPasswordChange(e) {
    this.setState({password: e.target.value})
  }
  onNewPasswordChange(e) {
    this.setState({newPassword: e.target.value})
  }

  onConfirmButtonClick() {
    if (this.state.userName === "") {
      this.setState({alertMsg: "username must not be empty"})
      return
    }
    if (this.state.newPassword === "") {
      this.setState({alertMsg: "new password must not be empty"})
      return
    }
    fetch(`/api/user/${this.state.userName}/password`, {
      method: "PUT",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        oldPassword: this.state.password,
        newPassword: this.state.newPassword,
      })
    }).then(async resp => {
      let res = await resp.json()
      if (resp.ok) {
        window.alert("Operation succeeded")
        this.props.navigate(-1)
      } else {
        this.setState({alertMsg: res.error})
        return
      }
    }).catch(e => {
      //todo: handle err
      console.log(e)
    })
  }

  render() {
    let message = null
    if (this.state.alertMsg !== "") {
      message = (<RetroCard danger style={{textAlign: "center", display: "block", margin: "20px auto", width: "60%"}}>{this.state.alertMsg}</RetroCard>)
    }

    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}
        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <RetroCard style={{width: "500px", margin: "30px auto", display: "block"}}>
            <div style={{margin: "50px", textAlign: "center"}}></div>
            <div style={{margin: "0 auto", textAlign: "center"}}>Admin Login: Please log in if you want to proceed</div>
            {message}
            <div style={{margin: "20px auto", textAlign: "center"}}>
              <span style={{fontSize: "20px"}}>Email{'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}</span>
              {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
              <RetroInput style={{fontSize: "20px"}} onChange={this.onUserNameChange} />
            </div>

            <div style={{margin: "20px auto", textAlign: "center"}}>
              <span style={{fontSize: "20px"}}>Old Password</span>

              <RetroInput password style={{fontSize: "20px"}} onChange={this.onPasswordChange} />
            </div>

            <div style={{margin: "20px auto", textAlign: "center"}}>
              <span style={{fontSize: "20px"}}>{'\u00a0'}New Password</span>

              <RetroInput password style={{fontSize: "20px"}} onChange={this.onNewPasswordChange} />
            </div>

            <RetroButton style={{margin: "20px auto", textAlign: "center", width: "50%", display: "block"}} onClick={this.onConfirmButtonClick}>Confirm</RetroButton>
            {/* <RetroButton style={{margin:"20px auto",textAlign:"center",width:"50%",display:"block"}}>Register</RetroButton> */}


          </RetroCard>
        </div>
      </div>
    )
  }
}
export default withRouter(ChangePassword)