import React, { useEffect, useState } from "react";
import {

Grid,
Card,
CardContent,
Typography,
CircularProgress,
Box,
} from "@mui/material";

const API_URL = "http://localhost:8081/events"; // Adjust port/path as needed

const Blockgrid = () => {
const [transactions, setTransactions] = useState([]);
const [loading, setLoading] = useState(true);

useEffect(() => {
    fetch(API_URL)
        .then((res) => res.json())
        .then((data) => {
            setTransactions(data);
            setLoading(false);
        })
        .catch(() => setLoading(false));
}, []);

if (loading) {
    return (
        <Box display="flex" justifyContent="center" mt={5}>
            <CircularProgress />
        </Box>
    );
}

return (
    <Box p={3}>
        <Typography variant="h4" gutterBottom>
            Transactions Dashboard
        </Typography>
        <Grid container spacing={3}>
            {transactions.map((tx) => (
                <Grid item xs={12} sm={6} md={4} lg={3} key={tx.id || tx.hash}>
                    <Card>
                        <CardContent>
                            <Typography variant="h6" gutterBottom>
                                Transaction #{tx.id || tx.hash}
                            </Typography>
                            <Typography variant="body2">
                                <strong>From:</strong> {tx.from}
                            </Typography>
                            <Typography variant="body2">
                                <strong>To:</strong> {tx.to}
                            </Typography>
                            <Typography variant="body2">
                                <strong>Amount:</strong> {tx.value}
                            </Typography>
                            {/* <Typography variant="body2">
                                <strong>Date:</strong> {tx.date}
                            </Typography> */}
                            {/* Add more fields as needed */}
                        </CardContent>
                    </Card>
                </Grid>
            ))}
        </Grid>
    </Box>
);
};

export default Blockgrid;