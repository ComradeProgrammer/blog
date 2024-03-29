import '../App.css';
import React from 'react';
import {withRouter} from './util';
import Header from '../components/Header';
import RetroCard from '../components/common/RetroCard';
import RetroInput from '../components/common/RetroInput';
import RetroButton from '../components/common/RetroButton';
import {Link} from "react-router-dom";
class LoginPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      userName: "",
      password: "",
      alertMsg: "",
    }
    this.onPasswordChange = this.onPasswordChange.bind(this)
    this.onUserNameChange = this.onUserNameChange.bind(this)
    this.onLoginButtonClick = this.onLoginButtonClick.bind(this)
  }

  onUserNameChange(e) {
    this.setState({userName: e.target.value})
  }
  onPasswordChange(e) {
    this.setState({password: e.target.value})
  }

  onLoginButtonClick() {
    fetch("/api/login", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        userName: this.state.userName,
        password: this.state.password,
      })
    }).then(async resp => {
      let res = await resp.json()
      if (resp.ok) {
        localStorage.setItem("user", JSON.stringify(res.user))
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
          <RetroCard style={{width: "500px", margin: "30px auto", height: "350px", display: "block"}}>
            <div style={{margin: "50px", textAlign: "center"}}></div>
            <div style={{margin: "0 auto", textAlign: "center"}}>User Login: Please log in if you want to proceed</div>
            {message}
            <div style={{margin: "20px auto", textAlign: "center"}}>
              <span style={{fontSize: "20px"}}>Email{'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}</span>
              {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
              <RetroInput style={{fontSize: "20px"}} onChange={this.onUserNameChange} />
            </div>

            <div style={{margin: "20px auto", textAlign: "center"}}>
              <span style={{fontSize: "20px"}}>Password</span>
              {'\u00a0\u00a0\u00a0\u00a0\u00a0\u00a0'}
              <RetroInput password style={{fontSize: "20px"}} onChange={this.onPasswordChange} />
            </div>

            <RetroButton style={{margin: "20px auto", textAlign: "center", width: "50%", display: "block"}} onClick={this.onLoginButtonClick}>Login</RetroButton>
            {/* <RetroButton style={{margin:"20px auto",textAlign:"center",width:"50%",display:"block"}}>Register</RetroButton> */}

            <Link to="/register" style={{color: "lightgreen", float: "right"}}>New User? Register an account</Link>

          </RetroCard>
        </div>
      </div>
    )
  }
}
export default withRouter(LoginPage)