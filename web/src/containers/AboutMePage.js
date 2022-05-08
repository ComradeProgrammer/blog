import '../App.css';
import React from 'react';
import Header from '../components/Header';
import Aboutme from '../assets/aboutme.md'
import RetroMarkDownDisplay from '../components/common/RetroMarkdownDisplay';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import {withRouter} from './util';

class AboutMePage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {}
    //fetch the aboutme.md
    fetch(Aboutme).then(res => {
      return res.text()
    }).then(res => {
      this.setState({aboutme: res})
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
            <RetroMarkDownDisplay markdown={this.state.aboutme}></RetroMarkDownDisplay>
          </RightBody>
        </div>
      </div>
    )
  }
}
export default withRouter(AboutMePage)