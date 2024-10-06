import { Link } from 'react-router-dom';
import img from '../assets/images/not-found.svg';
import Wrapper from '../assets/wrappers/ErrorPage';

export const Error = () => {
  return (
    <Wrapper className='full-page'>
    <div>
      <img src={img} alt='not found' />
      <h3>Page not found</h3>
      <p>Click below link and back to dashboard :D</p>
      <Link to='/'>back home</Link>
    </div>
  </Wrapper>
    
  )
}