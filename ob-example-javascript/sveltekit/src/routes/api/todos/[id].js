import { oceanbaseConnection } from '../../../lib/db/oceanbase';

export async function del(req) {
    const { id } = req.params;

    try {
        await oceanbaseConnection.query('DELETE FROM tasks WHERE id = ?', [id]);
        return {
            status: 200,
            body: { message: `Todo with id ${id} deleted successfully.` },
        };
    } catch (error) {
        console.error('Error deleting todo:', error);
        return {
            status: 500,
            body: { message: 'Error deleting todo. Please try again later.' },
        };
    }
}
export async function put(req) {
    const { id } = req.params;
    const { completed } = req.body;

    try {
        await oceanbaseConnection.query(
            'UPDATE tasks SET completed = ? WHERE id = ?',
            [completed, id]
        );
        const [updatedTodo] = await oceanbaseConnection.query(
            'SELECT * FROM tasks WHERE id = ?',
            [id]
        );
        return {
            status: 200,
            body: JSON.stringify(updatedTodo),
        };
    } catch (error) {
        console.error('Error toggling todo completed status:', error);
        return {
            status: 500,
            body: {
                message:
                    'Error toggling todo completed status. Please try again later.',
            },
        };
    }
}
