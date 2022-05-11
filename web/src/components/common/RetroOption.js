import React from 'react';

class RetroOption extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      style: {
        color: "lightgreen",
        backgroundColor: "black",
        border: "2px solid lightgreen",
      },
      data: props.data
    }
    for (let key in props.style) {
      this.state.style[key] = props.style[key]
    }
    this.onOptionChange=this.onOptionChange.bind(this)
  }

  UNSAFE_componentWillReceiveProps(nextProps) {
    let newStyle = {...this.state.style}
    for (let key in nextProps.style) {
      newStyle[key] = nextProps.style[key]
    }
    this.setState({style: newStyle, data: nextProps.data})
  }


  render(){
    let optionList=[]
    for(let i=0;i<this.state?.data?.length;i++){
      optionList.push((<option key={this.state.data[i].ID} value={this.state.data[i].ID}>{this.state.data[i].name}</option>))
    }
    return (
      <select style={this.state.style} onChange={this.onOptionChange}>
         {optionList}
      </select>
    )
  }

  onOptionChange(e){
    if(this.props.onChange){
      this.props.onChange(e)
    }
  }


}

export default RetroOption