import { writable } from 'svelte/store';

export const todos = writable([]);

export const loadTodos = async () => {
    try {
        const response = await fetch('/api/todos');
        const data = await response.json();
        todos.set(data);
    } catch (error) {
        console.error('Error fetching data:', error);
    }
};
loadTodos();

export const addTodo = async (text) => {
    try {
        const response = await fetch('/api/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ text: text, completed: false }),
        });

        const data = await response.json();
        todos.update((currentTodos) => [...currentTodos, data]);
    } catch (error) {
        console.error('Error adding todo:', error);
    }
};

export const deleteTodo = async (id) => {
    try {
        await fetch(`/api/todos/${id}`, {
            method: 'DELETE',
        });

        todos.update((currentTodos) =>
            currentTodos.filter((todo) => todo.id !== id)
        );
    } catch (error) {
        console.error('Error deleting todo:', error);
    }
};

export const toggleTodoCompleted = async (id, currentState) => {
    try {
        const response = await fetch(`/api/todos/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ completed: !currentState }),
        });

        const data = await response.json();
        todos.update((currentTodos) =>
            currentTodos.map((todo) => (todo.id === id ? data : todo))
        );
    } catch (error) {
        console.error('Error toggling todo completed status:', error);
    }
};
