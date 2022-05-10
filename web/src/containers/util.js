import {useLocation, useNavigate, useParams} from "react-router";
export function withRouter(Child) {
  return (props) => {
    const location = useLocation();
    const navigate = useNavigate();
    const params = useParams();
    //key here can trigger componentDidMount event when the router is moving backward
    return <Child {...props} navigate={navigate} location={location} params={params} key={location.pathname} />;
  }
}
