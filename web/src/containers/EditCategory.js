import '../App.css';
import React from 'react';
import Header from '../components/Header';
import LeftSideBar from '../components/LeftSideBar';
import RightBody from '../components/RightBody';
import {withRouter} from './util';
class EditCategory extends React.Component {
  constructor(props) {
    super(props);
    this.state = {}
  }

  render() {
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}

        <div style={{maxWidth: "1920px", minWidth: "755px", margin: "0 auto"}}>
          <LeftSideBar></LeftSideBar>
          <RightBody>

          </RightBody>
        </div>
      </div>
    )
  }
}
export default withRouter(EditCategory)