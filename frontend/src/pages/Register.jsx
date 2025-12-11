import { useState } from 'react';
import { register } from '../api/api';
import { useNavigate, Link } from 'react-router-dom'; // <-- CRITICAL FIX: Import Link

const Register = () => {
  const [form, setForm] = useState({ name: '', email: '', password: '' });
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      await register(form);
      alert('Registration Successful! Please Login.');
      navigate('/'); // Redirect to login page on success
    } catch (err) {
      // Improved error handling to safely check for response data
      const errorMessage = err.response?.data?.error || 'Something went wrong. Check console.';
      alert('Error: ' + errorMessage);
    } finally {
      setLoading(false);
    }
  };

  // --- STYLES ---
  const styles = {
    container: {
      minHeight: '100vh',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)', // Unique Gradient
      fontFamily: "'Segoe UI', Tahoma, Geneva, Verdana, sans-serif",
    },
    card: {
      background: 'rgba(255, 255, 255, 0.95)', // Slightly transparent white
      padding: '2.5rem',
      borderRadius: '20px',
      boxShadow: '0 20px 40px rgba(0, 0, 0, 0.2)', // Deep shadow for "pop"
      width: '100%',
      maxWidth: '400px',
      textAlign: 'center',
    },
    title: {
      color: '#333',
      marginBottom: '1.5rem',
      fontSize: '2rem',
      fontWeight: 'bold',
    },
    inputGroup: {
      marginBottom: '1.2rem',
      textAlign: 'left',
    },
    label: {
      display: 'block',
      marginBottom: '0.5rem',
      color: '#666',
      fontSize: '0.9rem',
      fontWeight: '600',
    },
    input: {
      width: '100%',
      padding: '12px 15px',
      borderRadius: '10px',
      border: '1px solid #ddd',
      fontSize: '1rem',
      outline: 'none',
      boxSizing: 'border-box', // Ensures padding doesn't break width
      transition: 'border-color 0.3s',
    },
    button: {
      width: '100%',
      padding: '14px',
      borderRadius: '10px',
      border: 'none',
      background: '#764ba2',
      color: 'white',
      fontSize: '1.1rem',
      fontWeight: 'bold',
      cursor: 'pointer',
      marginTop: '1rem',
      boxShadow: '0 4px 6px rgba(50, 50, 93, 0.11), 0 1px 3px rgba(0, 0, 0, 0.08)',
      transition: 'transform 0.2s',
    },
    linkText: {
      marginTop: '1.5rem',
      fontSize: '0.9rem',
      color: '#888',
    },
    link: {
      color: '#764ba2',
      textDecoration: 'none',
      fontWeight: 'bold',
      marginLeft: '5px',
    }
  };

  return (
    <div style={styles.container}>
      <div style={styles.card}>
        <h2 style={styles.title}>Create Account</h2>
        <form onSubmit={handleSubmit}>
          
          <div style={styles.inputGroup}>
            <label style={styles.label}>Full Name</label>
            <input 
              placeholder="e.g. Rahul Kale" 
              onChange={(e) => setForm({ ...form, name: e.target.value })} 
              required 
              style={styles.input} 
            />
          </div>

          <div style={styles.inputGroup}>
            <label style={styles.label}>Email Address</label>
            <input 
              placeholder="rahul@example.com" 
              type="email" 
              onChange={(e) => setForm({ ...form, email: e.target.value })} 
              required 
              style={styles.input} 
            />
          </div>

          <div style={styles.inputGroup}>
            <label style={styles.label}>Password</label>
            <input 
              placeholder="••••••••" 
              type="password" 
              onChange={(e) => setForm({ ...form, password: e.target.value })} 
              required 
              style={styles.input} 
            />
          </div>

          <button 
            type="submit" 
            style={styles.button}
            onMouseOver={(e) => e.target.style.transform = 'translateY(-2px)'}
            onMouseOut={(e) => e.target.style.transform = 'translateY(0)'}
            disabled={loading}
          >
            {loading ? 'Registering...' : 'Sign Up'}
          </button>
        </form>

        <div style={styles.linkText}>
          Already have an account? 
          <Link to="/" style={styles.link}>Login here</Link> {/* <-- CRITICAL FIX APPLIED HERE */}
        </div>
      </div>
    </div>
  );
};

export default Register;