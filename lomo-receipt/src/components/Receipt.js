import { useState } from 'react';
import QR from 'react-qr-code';
import Button from 'react-bootstrap/Button';

function Receipt(){
    const lnurl ="LNURL1DP68GURN8GHJ7MRWW4EXCTTSV9UJUMT99ACXZ7FLD46XW0TSV9UJVUPAVDSHYEPXV93KX0FNX5EN2VEEXVERXWFNXGENZVE3XVERXVPNX5ENXVECXVERXDFNXYN8V0F30HW4EG"
    const [copied, setCopied] = useState(false);

    const handleClick = () => {
        navigator.clipboard.writeText(lnurl);
        setCopied(true);
        setTimeout(() => setCopied(false), 1500); // reset copied state after 1.5 seconds
      };
    return (
        <>
        <QR value={lnurl}/>
        <br />
        <Button variant='light' 
        className='mt-4'
        onClick={handleClick}
        >  {copied ? 'Copied!' : 'Copy LNURL'}
        </Button>
        </>
    )
}

export default Receipt;