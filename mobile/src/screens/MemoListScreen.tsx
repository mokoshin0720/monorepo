import type { NativeStackScreenProps } from '@react-navigation/native-stack';
import { RootStackParamList } from 'App';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { View, StyleSheet, Text } from 'react-native';
import { Button } from 'src/components/Button';

type MemoListScreenNavigationProp = NativeStackScreenProps<
    RootStackParamList,
    'MemoList'
>;

export const MemoListScreen: React.FC<MemoListScreenNavigationProp> = ({
    navigation,
    route,
}) => {
    const [memos, setMemos] = useState<string[]>([]);

    useEffect(() => {
        getMemos();
    }, []);

    const getMemos = async () => {
        try {
            const response = await axios.get('http://localhost:3000/memos');
            console.log('response', response);
        } catch (error) {
            console.log(error);
        }
    };

    if (memos.length === 0) {
        return (
            <View style={emptyStyles.container}>
                <View style={emptyStyles.inner}>
                    <Text style={emptyStyles.title}>
                        最初のメモを作成しよう！
                    </Text>
                </View>
                <Button
                    label="作成する"
                    style={emptyStyles.button}
                    onPress={() => {
                        navigation.navigate('MemoCreate');
                    }}
                />
            </View>
        );
    }
};

const emptyStyles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
    inner: {
        justifyContent: 'center',
        alignItems: 'center',
    },
    title: {
        fontSize: 18,
        marginBottom: 24,
    },
    button: {
        alignSelf: 'center',
    },
});
