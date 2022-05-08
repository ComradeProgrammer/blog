import React from 'react';
import Logo from "../assets/logo.png"
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
      currentUserName: null,
      currentUserID: null,
    }
  }

  render() {
    return (
      <div style={this.state.style} >
        <img style={{height: "60px", width: "60px", display: "inline-block"}} src={Logo} alt="" />

        <div style={{width: "20%", display: "inline-block"}}>A Blog in the style of 1980s</div>

        <RetroButton style={{fontSize: "25px", width: "15%"}}>Blogs</RetroButton>
        <RetroButton style={{fontSize: "25px", width: "15%"}}>About Me</RetroButton>
        <RetroButton style={{fontSize: "25px", width: "15%"}}>Resume/CV</RetroButton>
        <RetroButton style={{fontSize: "25px", width: "15%"}}>To Github</RetroButton>

        <RetroButton style={{float: "right", height: "62px", width: "62px", lineHeight: "62px", textAlign: "center"}}>{this.state.currentUserName === null ? "Login" : this.state.currentUserName}</RetroButton>
      </div>
    );
  }
}
export default Header;