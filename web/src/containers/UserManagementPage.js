import '../App.css';
import React from 'react';
import {withRouter, isAdmin} from './util';
import Header from '../components/Header';
import RetroCard from '../components/common/RetroCard';
import RetroButton from '../components/common/RetroButton';


class UserManagementPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      users: [],
    }
    this.onDeleteUserClick = this.onDeleteUserClick.bind(this)
  }

  componentDidMount() {
    fetch("/api/user").then(res => {
      //todo: handle res code
      return res.json()
    }).then((res) => {
      this.setState({users: res})
    }).catch(e => {
      //todo: handle err
      console.log(e)
    })
  }

  onDeleteUserClick(id) {
    let result = window.confirm("Do you really want to delete this user?")
    if (result) {
      fetch(`/api/user/${id}`, {
        method: "DELETE"
      }).then(async res => {
        if (res.ok) {
          alert("Operation succeeded")
          window.location.reload()
        } else {
          alert("Operation Failed")
          let body = await res.json()
          console.log(body)
        }
      }).catch(e => {
        console.log(e)
      })
    }
  }


  generateTable() {
    let admin = isAdmin()
    let rows = []
    rows.push((
      <tr key="header">
        <th>ID</th>
        <th>UserName</th>
        <th>IsAdmin</th>
        <th>Created at</th>
        <th>Operations</th>
      </tr>
    ))
    for (let i = 0; i < this.state.users.length; i++) {
      rows.push((
        <tr key={this.state.users[i].ID}>
          <th>{this.state.users[i].ID}</th>
          <th>{this.state.users[i].userName}</th>
          <th>{this.state.users[i].isAdmin ? "true" : "false"}</th>
          <th>{this.state.users[i].createAt}</th>
          {
            admin ? (
              <th>
                <RetroButton danger style={{margin: "5px"}} onClick={() => {this.onDeleteUserClick(this.state.users[i].ID)}}>Delete User</RetroButton>
              </th>
            ) : (<th></th>)
          }
        </tr>
      ))
    }
    return rows
  }

  render() {
    return (
      <div className='body'>
        <Header></Header>
        <div style={{height: "80px"}}></div>{/* place holder for header */}

        <RetroCard style={{width: "80%", minWidth: "755px", margin: "0 auto", minHeight: "350px", display: "block"}}>
          <div style={{margin: "0 auto", textAlign: "center", fontSize: "30px"}}>All Users</div>
          <table border="1" style={{border: "1px solid lightgreen", margin: "0 auto", fontSize: "20px"}}>
            <tbody>
              {this.generateTable()}
            </tbody>
          </table>
        </RetroCard>
      </div>
    )
  }
}

export default withRouter(UserManagementPage)
