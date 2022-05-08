import React from 'react';
import RetroCard from '../components/common/RetroCard';
import RetroButton from './common/RetroButton';

class LeftSideBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      categories: []
    }
  }

  componentDidMount() {
    fetch("/api/category").then(res => {
      //todo: handle res code
      return res.json()
    }).then((res) => {
      this.setState({categories: res})
    }).catch(e => {
      //todo: handle err
      console.log(e)
    })
  }

  render() {
    let categoryButtons = []
    for (let i = 0; i < this.state.categories.length; i++) {
      categoryButtons.push(
        (<RetroButton key={this.state.categories[i].ID} style={{width: "90%", display: "block", margin: "3px auto"}}>{this.state.categories[i].name}</RetroButton>)
      )
    }
    return (
      <RetroCard style={{width: "20%", height: "100px", verticalAlign: "top", textAlign: "center"}}>
        All Blog categories<br />
        {categoryButtons}
      </RetroCard>
    )
  }
}
export default LeftSideBar