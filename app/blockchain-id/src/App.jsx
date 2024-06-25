import Navbar from './components/Navbar'
import Home from './components/Home'
import Features from './components/Features'
import IssueIdentity from './components/IssueIdentity'

const App = () => {
  return (
    <>
      <Navbar />
      <div className="max-w-7xl mx-auto pt-20 px-6">
          <Home />
          <Features />
          <IssueIdentity />
      </div>
    </>
  )
}

export default App