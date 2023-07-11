import { StyleSheet } from 'react-native';

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { MemoListScreen } from 'src/screens/MemoListScreen';
import { MemoCreateScreen } from 'src/screens/MemoCreateScreen';

export type RootStackParamList = {
    MemoList: undefined;
    MemoCreate: undefined;
};

const Stack = createNativeStackNavigator<RootStackParamList>();

export default function App() {
    return (
        <NavigationContainer>
            <Stack.Navigator>
                <Stack.Screen name="MemoList" component={MemoListScreen} />
                <Stack.Screen name="MemoCreate" component={MemoCreateScreen} />
            </Stack.Navigator>
        </NavigationContainer>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
});
