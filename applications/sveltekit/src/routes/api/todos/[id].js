import { createOceanbaseConnection } from '../../../lib/db';

export async function del(req) {
  const { id } = req.params;

  try {
    await createOceanbaseConnection().query('DELETE FROM tasks WHERE id = ?', [
      id,
    ]);
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
    await createOceanbaseConnection().query(
      'UPDATE tasks SET completed = ? WHERE id = ?',
      [completed, id],
    );
    const [updatedTodo] = await createOceanbaseConnection().query(
      'SELECT * FROM tasks WHERE id = ?',
      [id],
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
