import React from 'react';
import Logo from "../assets/logo.png"
import {withRouter} from '../containers/util';
import RetroButton from './common/RetroButton';

class Header extends React.Component {
  constructor(props) {
    super(props);
    let color = "lightgreen"
    //set style
    this.state = {
      style: {
        height: "70px",
        width: "100%",
        color: color,
        backgroundColor: "black",
        borderBottom: "1px solid " + color,
        position: "fixed",
        top: "0px",
        left: "0px",
        margin: 0,
      },
      user: null
    }
    this.onBlogsButtonClick = this.onBlogsButtonClick.bind(this)
    this.onAboutMeButtonClick = this.onAboutMeButtonClick.bind(this)
    this.onLoginButtonClick=this.onLoginButtonClick.bind(this)
  }

  componentDidUpdate() {
    if (sessionStorage.getItem("user")) {
      try {
        let user = JSON.parse(sessionStorage.getItem("user"))
        if (user.userName) {
          if (user.userName !== this.state?.user?.userName) {
            this.setState({user: user})
          }
        } else {
          throw Error("invalid session")
        }
      } catch (e) {
        sessionStorage.removeItem("user")
        this.setState({user: null})
      }
    }
  }

  render() {

    console.log(this.state.user == null ?"62px":"0")
    return (
      <div style={this.state.style} >
        <img style={{height: "60px", width: "60px", display: "inline-block"}} src={Logo} alt="" />

        <div style={{width: "20%", display: "inline-block"}}>A Blog in the style of 1980s</div>

        <RetroButton style={{fontSize: "25px", width: "15%"}} onClick={this.onBlogsButtonClick}>Blogs </RetroButton>
        <RetroButton style={{fontSize: "25px", width: "15%"}} onClick={this.onAboutMeButtonClick}>About Me</RetroButton>
        <RetroButton style={{fontSize: "25px", width: "15%"}} >Resume/CV</RetroButton>
        <RetroButton style={{fontSize: "25px", width: "15%"}} onClick={this.onGithubButtonClick}>To Github</RetroButton>

        <RetroButton style={{float: "right", height: "62px", minWidth: "62px",lineHeight: this.state.user == null ?"62px":"0",  textAlign: "center"}} onClick={this.onLoginButtonClick}>
          {this.state.user === null ? "Login" : this.state.user.userName+" Logout" } 
        </RetroButton>
      </div>
    );
  }

  onBlogsButtonClick() {
    this.props.navigate("/blog")
  }

  onAboutMeButtonClick() {
    this.props.navigate("/aboutme")
  }

  onGithubButtonClick() {
    window.location.href = "https://github.com/ComradeProgrammer"
  }

  onLoginButtonClick(){
    if(this.state.user!==null){
      //todo: use logout api
      sessionStorage.removeItem("user")
      this.setState({user: null})
    }
    this.props.navigate("/login")
  }

}
export default withRouter(Header);