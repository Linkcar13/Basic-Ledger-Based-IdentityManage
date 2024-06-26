import Navbar from './components/Navbar'
import Home from './components/Home'
import Features from './components/Features'
import IssueIdentity from './components/IssueIdentity'
import ValidateIdentity from './components/ValidateIdentity'
import RevokeIdentity from './components/RevokeIdentity'
import Footer from './components/Footer'

const App = () => {
  return (
    <>
      <Navbar />
      <div className="max-w-7xl mx-auto pt-20 px-6">
          <Home />
          <Features />
          <IssueIdentity />
          <ValidateIdentity />
          <RevokeIdentity />
          <Footer />
      </div>
    </>
  )
}

export default App