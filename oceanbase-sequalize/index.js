const express = require('express');
const app = express();

// Connect to the database
const db = require('./models');
const { Contact } = require('./models');

// Middleware to parse JSON bodies
app.use(express.json());

app.get('/contacts', async (req, res) => {
    // Get all contacts
    const contacts = await Contact.findAll();
    res.json(contacts);
});

app.get('/contact/:id', async (req, res) => {
    // Get a contact by id
    const contact = await Contact.findByPk(req.params.id);
    if (contact) {
        res.json(contact);
    } else {
        res.status(404).send({
            status: 'failed',
            message: 'Contact not found',
        });
    }
});

app.post('/contact', async (req, res) => {
    // Create a new contact
    const newContact = await Contact.create(req.body);
    res.json(newContact);
});

app.delete('/contact/:id', async (req, res) => {
    // Delete a contact by id
    const result = await Contact.destroy({
        where: {
            id: req.params.id,
        },
    });
    if (result) {
        res.status(200).send({
            status: 'succeed',
            message: 'Contact deleted',
        });
    } else {
        res.status(404).send({
            status: 'failed',
            message: 'Contact not found',
        });
    }
});

app.patch('/contact/:id', async (req, res) => {
    // Update a contact by id
    const updatedContact = await Contact.update(req.body, {
        where: {
            id: req.params.id,
        },
    });
    console.log(updatedContact);
    if (updatedContact[0]) {
        res.status(200).send({
            status: 'Contact updated',
            data: await Contact.findByPk(req.params.id),
        });
    } else {
        res.status(404).send({
            status: 'failed',
            message: 'Contact not found',
        });
    }
});

db.sequelize.sync().then((req) => {
    app.listen(3000, () => {
        console.log('Server running at port 3000...');
    });
});
