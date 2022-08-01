import '../App.css';
import React from 'react';
import {withRouter, getUser} from './util';
import Header from '../components/Header';
import RetroCard from '../components/common/RetroCard';
import RetroButton from '../components/common/RetroButton';

class UserPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {

    }
  }

  onLogoutButtonClick() {
    let user = getUser()
    if (user !== null) {
      //todo: use logout api
      fetch("/api/login", {
        method: "DELETE",
      })
      localStorage.removeItem("user")
    }
    this.props.navigate(-1)
  }

  onChangePasswordButtonClick() {
    this.props.navigate("/changepassword")

  }

  render() {
    let user = getUser()
    if (user === null) {
      this.props.navigate("/login")
    }
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}

        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <RetroCard style={{width: "50%", minWidth: "755px", margin: "0 auto", minHeight: "350px", display: "block"}}>
            <div style={{margin: "50px", textAlign: "center"}}></div>
            <div style={{margin: "0 auto", textAlign: "center", fontSize: "30px"}}>User Settings</div>
            <div style={{margin: "0 auto", textAlign: "center"}}>Greetings, user {user.userName}</div>


            <RetroButton warning style={{margin: "20px auto", textAlign: "center", width: "50%", display: "block"}} onClick={this.onLogoutButtonClick.bind(this)}>Log out</RetroButton>
            <RetroButton style={{margin: "20px auto", textAlign: "center", width: "50%", display: "block"}} onClick={this.onChangePasswordButtonClick.bind(this)}>Change Password</RetroButton>



          </RetroCard>
        </div>

        <div style={{height: "80px"}}></div>{/* place holder for header */}
      </div>
    )
  }
}

export default withRouter(UserPage)